#!/bin/bash

set -e

function gen_dao() {
    table_name=$1
    IFS='_'
    read -ra arr <<< "$table_name"
    unset IFS

    MYSQL_WHITELIST="$table_name" sqlboiler -c sqlboiler_dao.toml --no-driver-templates -d --templates pkg/data/templates/sqlboiler -p dao --output "pkg/dao" mysql

    struct_name=$(gsed -nE 's/^type (\w+) struct \{/\1/p' "pkg/dao/impl/$table_name.go")
    cmd=(ifacemaker)
    cmd+=(-f "pkg/dao/impl/${table_name}.go")

    if [ -f "pkg/dao/impl/${table_name}.extend.go" ]; then
        cmd+=(-f "pkg/dao/impl/${table_name}.extend.go")
    fi
    cmd+=( -s "${struct_name}" -i "${struct_name}" -p dao -o "pkg/dao/${table_name}.go")

    echo "${cmd[@]}"
    "${cmd[@]}"

    goimports -w "pkg/dao/${table_name}.go" "pkg/dao/impl/${table_name}.go"
    if [ -f "pkg/dao/impl/${table_name}.extend.go" ]; then
        goimports -w "pkg/dao/impl/${table_name}.extend.go"
    fi
}

# sqlboiler related environment
source ./sqlboiler.env

echo "generating models for: $MYSQL_WHITELIST"
sqlboiler --wipe -d --no-tests --templates "$DRIVER_TEMPLATE_DIR" --output pkg/dao/models mysql

read -ra tables <<< "$MYSQL_WHITELIST"
for table in "${tables[@]}"
do
  echo "generating dao code for: $table"
  gen_dao "$table"
done
