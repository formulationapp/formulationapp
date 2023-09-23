export default interface Block {
    type: string
    data: {
        label: string
        text: string
        level: number
    }
}