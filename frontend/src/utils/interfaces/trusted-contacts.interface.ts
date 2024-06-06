export interface ITrustedContact {
    id: number
    value: string
    label: string
    domain: string
    type: string
    addedBy?: {
        id: string
        name: string
    }
}