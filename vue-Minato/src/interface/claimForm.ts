export interface ClaimForm {
    claimid: string
    username: string
    userid: string
    callnumber: string
    insuranceid: number | null
    carid: string
    region: string
    date: string
    claimfile: string[]
}
export interface ClaimResult {
    claimid: string
    username: string
    userid: string
    callnumber: string
    insuranceid: number | null
    carid: string
    region: string
    date: string
    status: number
    compensation: number
}
