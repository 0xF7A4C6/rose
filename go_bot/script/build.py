import os, threading, time

arch = [
        #'darwin:386',
        #'darwin:amd64',
        #'dragonfly:amd64',
        #'freebsd:386',
        #'freebsd:amd64',
        #'freebsd:arm',
        'linux:386',
        'linux:amd64',
        'linux:arm',
        'linux:arm64',
        'linux:ppc64',
        'linux:ppc64le',
        'linux:mips',
        'linux:mipsle',
        'linux:mips64',
        'linux:mips64le',
        #'netbsd:386',
        #'netbsd:amd64',
        #'netbsd:arm',
        #'openbsd:386',
        #'openbsd:amd64',
        #'openbsd:arm',
        #'plan9:386',
        #'plan9:amd64',
        #'solaris:amd64',
]

def cross_compile():
    def compile(go_arch: str, go_os: str):
        os.system(f'GOOS={go_os} GOARCH={go_arch} go build -a -gcflags=all="-l -B" -ldflags="-w -s" -o ./bin/builds/rose_{go_arch}{".exe" if go_os == "windows" else ""} ../ && ./upx -9 bin/builds/rose_{go_arch}{".exe" if go_os == "windows" else ""}')
        print(f'* {go_os} --> {go_arch}')

    for proc in arch:
        parsed = proc.split(':')
        go_os, go_arch = parsed[0], parsed[1]

        while threading.active_count() >= 5:
            time.sleep(0.5)
        
        threading.Thread(target=compile, args= [go_arch, go_os]).start()

os.system('rm -rf ./bin && mkdir ./bin && mkdir ./bin/builds && clear')
print('''
>> Rose cross compiler + packer.
''')
cross_compile()