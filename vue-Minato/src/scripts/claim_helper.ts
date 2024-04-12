//跟Minatosys.ts一起辅助完成交易流程
import axios from "axios"
export async function updateAddress(address: string, claimid: string) {
    let updateaddress = {
        claim_id: claimid,
        address: address,
    }
    await axios
        .post("http://localhost:8080/updateaddress", updateaddress)
        .then((response) => {
            console.log(response.data)
        })
}
export async function updateCompensation(
    compensation: number,
    claimid: string,
) {
    let updatestatus = {
        claim_id: claimid,
        compensation: compensation,
    }
    await axios
        .post("http://localhost:8080/updatestatus", updatestatus)
        .then((response) => {
            console.log(response.data)
        })
}
