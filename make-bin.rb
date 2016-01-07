OUTPUT_FILE = 'testbin'

prog = [
  9, 32768, 32769,
  4, 19, 32768
]

output = open(OUTPUT_FILE, 'w+')

# use prog.pack to generate the bytes
# 's<*' means 'little endian integers'
# http://ruby-doc.org/core-2.2.0/Array.html#method-i-pack
output.puts(prog.pack('s<*'))

output.close

