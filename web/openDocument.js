var child_process = require('child_process')

var url = 'https://tome.app/bal-248/clpjv0pla0mwepq7db4fvmbdf'
var cmd = ''
console.log(process.platform)
switch (process.platform) {
  case 'win32':
    cmd = 'start'
    child_process.exec(cmd + ' ' + url)
    break

  case 'darwin':
    cmd = 'open'
    child_process.exec(cmd + ' ' + url)
    break
}
