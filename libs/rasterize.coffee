page = require('webpage').create()
system = require 'system'

if system.args.length < 3 or system.args.length > 4
  console.log 'Usage: rasterize.coffee URL output [paperwidth*paperheight|paperformat]'
  console.log '  output examples: out.pdf /dev/stdout'
  console.log '  paper (pdf output) examples: "5in*7.5in", "10cm*20cm", "A4", "Letter"'
  phantom.exit 1

else
  address = system.args[1]
  output = system.args[2]
  format = output.split('.').pop()
  if format not in ['png', 'pdf', '/dev/stdout']
    console.log 'Invalid output format - filename should end in png or pdf'
    phantom.exit()

  if format is '/dev/stdout'
    format = 'pdf'

  # If a paper size is specified
  if system.args.length is 4
    size = system.args[3].split '*'
    # Specified something like 5in*7.5in
    if size.length is 2
      page.paperSize = { width: size[0], height: size[1], border: '0px' }
    # Otherwise single value (i.e. Letter)
    else
      page.paperSize = { format: system.args[3], orientation: 'portrait', border: '1cm' }

  page.open address, (status) ->
    if status isnt 'success'
      console.log 'Unable to load the address!'
      phantom.exit()
    else
      # We force PDF output
      window.setTimeout (-> page.render(output, {format: format}); phantom.exit()), 200
