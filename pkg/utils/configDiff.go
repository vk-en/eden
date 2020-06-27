package utils

var defaultEnvDiffConfig = `#config is generated by eden
adam:
    #tag on adam container to pull
    tag: {{ .DefaultAdamTag }}

    #port of adam
    port: {{ .DefaultAdamPort }}

    #domain of adam
    domain: {{ .DefaultDomain }}

    #ip of adam for EVE access
    eve-ip: {{ .IP }}

    #ip of adam for EDEN access
    ip: {{ .IP }}

    #certificate for communication with adam
    ca: {{ .DefaultAdamDist }}/run/config/root-certificate.pem

    redis:
      #host of adam's redis for EDEN access
      eden: redis://{{ .IP }}:{{ .DefaultRedisPort }}
      #host of adam's redis for ADAM access
      adam: redis://{{ .DefaultRedisContainerName }}:{{ .DefaultRedisPort }}

eve:
    #live image of EVE
    image-file: {{ .DefaultImageDist }}/eve/live.img

    #devmodel
    devmodel: ZedVirtual-4G

    #file to save qemu config
    qemu-config: {{ .EdenDir }}/qemu.conf

    #EVE arch (amd64/arm64)
    arch: {{ .Arch }}

    #EVE acceleration (set to false if you have problems with qemu)
    accel: true

    #uuid of EVE to use in cert
    uuid: {{ .UUID }}

    #serial number in SMBIOS
    serial: "{{ .DefaultEVESerial }}"

    #onboarding certificate of EVE to put into adam
    cert: certs/onboard.cert.pem

    #EVE firmware
    firmware: [{{ .DefaultImageDist }}/eve/OVMF_CODE.fd,{{ .DefaultImageDist }}/eve/OVMF_VARS.fd]

    #forward of ports in qemu [(HOST:EVE)]
    hostfwd:
        {{ .DefaultSSHPort }}: 22
        5912: 5901
        5911: 5900
        8027: 8027
        8028: 8028

eden:
    #root directory of eden
    root: {{ .Root }}

    #download eve instead of build
    download: true

    #eserver is tool for serve images
    eserver:
        #ip (domain name) of eserver for EVE access
        ip: {{ .DefaultDomain }}

        #port for eserver
        port: {{ .DefaultEserverPort }}

    #ssh-key to put into EVE
    ssh-key: {{ .DefaultSSHKey }}

    #test binary
    test-bin: "{{ .DefaultTestProg }}"

    #test scenario
    test-scenario: "{{ .DefaultTestScenario }}"

redis:
    #port for access redis
    port: {{ .DefaultRedisPort }}

    #tag for redis image
    tag: {{ .DefaultRedisTag }}
`

//GenerateConfigFileDiff is a function to generate diff yml for new context
func GenerateConfigFileDiff(filePath string) error {
	return generateConfigFileFromTemplate(filePath, defaultEnvDiffConfig)
}
