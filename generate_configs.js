var fs = require('fs');

var size = 2;
var template = fs.readFileSync('cell/cell.yaml');
var files = [];
for (var i=0;i<size;i++) {
    for (var j=0;j<size;j++) {
        var name = 'cell/cell'+i+'x'+j+".yaml";
        var c = template.toString().replace('$NAME$','cell'+i+'x'+j);
        fs.writeFileSync(name,c);
        files.push(name);
    }
}
fs.writeFileSync('start.sh',"goapp serve frontend/app.yaml "+files.join(' '));
fs.writeFileSync('deploy.sh',"goapp deploy frontend/app.yaml "+files.join(' '));
