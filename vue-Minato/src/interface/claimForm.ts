export interface ClaimForm {
    username: string
    userid: string
    callnumber: string
    insuranceid: string
    carid: string
    region: string
    date: string
    claimfile: string
}
export interface ClaimResult {
    claimid: string
    username: string
    userid: string
    callnumber: string
    insuranceid: string
    carid: string
    region: string
    date: string
    status: number
    compensation: number
}
