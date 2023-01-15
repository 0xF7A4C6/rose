import pyasn

asndb = pyasn.pyasn('as.db')
ranges = []

for ip in list(set(open('./ips.txt', 'r+').read().splitlines())):
    asn, prefix = asndb.lookup(ip)
    print(asn, prefix)
    
    with open('./prefix.txt', 'a+') as pf:
        pf.write(prefix + '\n')
    
    ranges += list(asndb.get_as_prefixes(asn))

ranges = list(set(ranges))
with open('./r.txt', 'a+') as f:
    for r in ranges:
        f.write(f'{r}\n')