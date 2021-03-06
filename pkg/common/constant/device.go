package constant

const DeviceMobileConfig = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">

<plist version="1.0">
    <dict>
        <key>PayloadContent</key>
        <dict>
            <key>URL</key>
            <string>%s</string>
            <key>DeviceAttributes</key>
            <array>
                <string>UDID</string>
                <string>IMEI</string>
                <string>ICCID</string>
                <string>VERSION</string>
                <string>PRODUCT</string>
            </array>
        </dict>
        <key>PayloadOrganization</key>
        <string>www.dumpapp.com</string>
        <key>PayloadDisplayName</key>
        <string>Dumpapp</string>
        <key>PayloadVersion</key>
        <integer>1</integer>
        <key>PayloadUUID</key>
        <string>8C7AD0B8-3900-44DF-A52F-3C4F92921807</string>
        <key>PayloadIdentifier</key>
        <string>com.yun-bangshou.profile-service</string>
        <key>PayloadDescription</key>
        <string>该配置文件将帮助用户获取当前iOS设备的UDID号码。This temporary profile will be used to find and display your current device's UDID.</string>
        <key>PayloadType</key>
        <string>Profile Service</string>
    </dict>
</plist>
`
