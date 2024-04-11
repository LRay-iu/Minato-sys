//这个文档里写的内容主要是有关跟合约交互的部分，
// 将合约交互的方法从项目之中剥离出来，
// 可以优化scripts中代码过于冗余的问题
import { ethers } from "@/scripts/ethers-5.7.esm.min.js"
import { abi, contractAddress } from "./constants"

const provider = new ethers.providers.Web3Provider((window as any).ethereum)
const signer = provider.getSigner()
//创建智能合约实例
const contract = new ethers.Contract(contractAddress, abi, signer)

// 用以检查浏览器有没有安装Metamask扩展
async function connect() {
    if (typeof (window as any).ethereum !== "undefined") {
        await (window as any).ethereum.request({
            method: "eth_requestAccounts",
        })
    } else {
        console.log("No metamask!")
    }
}
//管理员转账
async function compensation(toAddress: any, amountInWei: BigInt) {
    console.log("Funding with", toAddress)
    if (typeof (window as any).ethereum !== "undefined") {
        const ownerAddress = await contract.i_owner()
        const tx = await contract.withdrawToAddress(
            toAddress,
            ethers.BigNumber.from(amountInWei),
        )
        // const tx = await contract.withdraw()
        await tx.wait()
        console.log("Transaction successful")
    } else {
        console.log("No metamask!")
    }
}
//用户支付保险费用
async function fund(toAddress: any, amountInWei: BigInt) {
    console.log("Funding with", toAddress)
    if (typeof (window as any).ethereum !== "undefined") {
        const ownerAddress = await contract.i_owner()
        const tx = await contract.fund(ethers.BigNumber.from(amountInWei))
        await tx.wait()
        console.log("Transaction successful")
    } else {
        console.log("No metamask!")
    }
}

export { connect, compensation }
