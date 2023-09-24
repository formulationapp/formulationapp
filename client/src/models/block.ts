export default interface Block {
    id: string
    type: string
    data: {
        label: string
        text: string
        level: number
    }
}