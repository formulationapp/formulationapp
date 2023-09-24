export function markdownToHtml(mdString: string) {
    // Adapted from https://randyperkins2k.medium.com/writing-a-simple-markdown-parser-using-javascript-1f2e9449a558
    return mdString
        .replace(/\\\*/g, '\\*\\*')
        .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
        .replace(/\*(.*?)\*/g, '<em>$1</em>')
        .replaceAll('\\<em>\\</em>', '*')
        // .replace(/(?<!\\)\*(?<!\\)\*(.*?)(?<!\\)\*(?<!\\)\*/g, '<strong>$1</strong>')
        // .replace(/(?<!\\)\*(.*?)(?<!\\)\*/g, '<em>$1</em>')
        // .replace(/\\\*/g, '*')
        .trim()
}

export function htmlToMarkdown(htmlString: string) {
    return htmlString
        .replaceAll('<p>', '')
        .replaceAll('</p>', '')
        .replaceAll('<strong>', '**')
        .replaceAll('</strong>', '**')
        .replaceAll('<em>', '*')
        .replaceAll('</em>', '*')
        .replaceAll(/\<br.*?\>/g, '')
}