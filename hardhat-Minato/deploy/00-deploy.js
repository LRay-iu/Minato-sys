//00-deploy-mocks.js
//这段是部署本地预言机
const { network } = require("hardhat");

const DECIMALS = "8";
const INITIAL_PRICE = "200000000000"; // 2000

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const chainId = network.config.chainId;
  log(network.name);
  if (chainId == 31337) {
    log("Local network detected!Deploying mocks...");
    await deploy("MockV3Aggregator", {
      contract: "MockV3Aggregator",
      from: deployer,
      log: true,
      args: [DECIMALS, INITIAL_PRICE],
    });
    log("Mocks deployed!");
    log("--------------------------------------------------------");
  }
};
// 这段代码指定了当前部署脚本相关的标签。在这里，使用了两个标签："all" 和 "mocks"。
// "all" 标签： 这个标签可能用于将部署脚本与整个项目的所有部署任务关联起来。
// 当运行 npx hardhat deploy --tags all 时，将运行所有带有 "all" 标签的部署任务。
// "mocks" 标签： 这个标签可能用于将部署脚本与与模拟合约相关的其他部署任务关联起来。
// 当运行 npx hardhat deploy --tags mocks 时，将运行所有带有 "mocks" 标签的部署任务。
//yarn hardhat deploy --tags mocks
module.exports.tags = ["all", "mocks"];