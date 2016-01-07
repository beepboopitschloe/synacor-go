OUTPUT_FILE = 'testbin'

prog = [
  9, 32768, 72,
  19,
  9, 32768, 101,
  19,
  9, 32768, 108,
  19,
  9, 32768, 108,
  19,
  9, 32768, 111,
  19,
  9, 32768, 44,
  19,
  9, 32768, 32,
  19,
  9, 32768, 119,
  19,
  9, 32768, 111,
  19,
  9, 32768, 114,
  19,
  9, 32768, 108,
  19,
  9, 32768, 100,
  19,
  9, 32768, 33,
  19,
  9, 32768, 13,
  19,
  0
]

output = open(OUTPUT_FILE, 'w+')

# use prog.pack to generate the bytes
# 's<*' means 'little endian integers'
# http://ruby-doc.org/core-2.2.0/Array.html#method-i-pack
output.puts(prog.pack('s<*'))

output.close

