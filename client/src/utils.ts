import { createToaster } from "@meforma/vue-toaster";

export function capitalize(s: string): string {
    return s[0].toUpperCase() + s.slice(1);
}

Array.prototype.swap = function (x, y) {
    if (x < 0 || y < 0) return this;
    if (x >= this.length || y >= this.length) return this;
    const b = this[x];
    this[x] = this[y];
    this[y] = b;
    return this;
}
export async function copyToClipboard(text) {
    navigator.clipboard.writeText(text).then(function() {
        const toaster = createToaster({ /* options */ });
        toaster.success('Copied to clipboard!');
    }, function(err) {
        window.open(link.value, '_blank');
    });
}