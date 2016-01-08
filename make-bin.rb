OUTPUT_FILE = 'testbin'

prog = [
  1, 32768, 72,
  19, 32768,
  1, 32768, 101,
  19, 32768,
  1, 32768, 108,
  19, 32768,
  1, 32768, 108,
  19, 32768,
  1, 32768, 111,
  19, 32768,
  1, 32768, 44,
  19, 32768,
  1, 32768, 32,
  19, 32768,
  1, 32768, 119,
  19, 32768,
  1, 32768, 111,
  19, 32768,
  1, 32768, 114,
  19, 32768,
  1, 32768, 108,
  19, 32768,
  1, 32768, 100,
  19, 32768,
  1, 32768, 33,
  19, 32768,
  1, 32768, 10,
  19, 32768,
  0
]

output = open(OUTPUT_FILE, 'w+')

# use prog.pack to generate the bytes
# 's<*' means 'little endian integers'
# http://ruby-doc.org/core-2.2.0/Array.html#method-i-pack
output.puts(prog.pack('s<*'))

output.close

