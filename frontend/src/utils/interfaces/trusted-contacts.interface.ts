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
    type: CONTACT
    addedBy?: {
        id: string
        name: string
    }
}