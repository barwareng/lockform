export enum CONTACT {
    EMAIL = 'email',
    PHONE = 'phone',
    ADDRESS = 'address'
}
export interface IContact {
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