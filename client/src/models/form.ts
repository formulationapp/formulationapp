import type Data from "@/models/data";

export default interface Form {
    ID: number,
    name: string
    data: Data
    workspaceID: number
    secret: string
    UpdatedAt: string
    CreatedAt: string
}