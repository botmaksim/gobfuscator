const fs = require('fs');

function main() {
    try {
        const input = fs.readFileSync('/dev/stdin', 'utf-8').trim().split(/\s+/);
        if (input.length === 0 || input[0] === '') return;
        const n = parseInt(input[0], 10);
        for (let i = 1; i <= n; i++) {
            const s = input[i];
            if (s && s.length > 10) {
                console.log(s[0] + (s.length - 2) + s[s.length - 1]);
            } else if (s) {
                console.log(s);
            }
        }
    } catch (e) {}
}

main();
