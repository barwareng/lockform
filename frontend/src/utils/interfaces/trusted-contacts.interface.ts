export enum CONTACT {
    EMAIL = 'email',
    PHONE = 'phone',
    ADDRESS = 'address'
}
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