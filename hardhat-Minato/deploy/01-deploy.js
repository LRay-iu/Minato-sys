//01-deploy-fund-me.js
//helpconfig代表了helper-hardhat-config.js这个文件
// const helperconfig = require("../helper-hardhat-config.js")
// const networkconfig = helperconfig.networkConfig
//node.js的快捷写法，写法等同于上方的
const {
    networkConfig,
    developmentChains,
  } = require("../helper-hardhat-config.js");
  const { network } = require("hardhat");
  const { verify } = require("../utils/verify.js");
  //async function deployFunc(hre) {
  //     hre.getNameAccounts()
  //     hre.deployments
  // }
  // module.exports.default = deployFunc //将deplyFunc设置为默认要找的函数
  
  //hre代表hardhat运行环境
  //写法等同于上方的
  // module.exports = async (hre) => {
  //     const {getNameAccounts,deployments} = hre
  // }
  //node.js的语法糖，写法等同于上方的
  module.exports = async ({ getNamedAccounts, deployments }) => {
    //将deploy和log从deployments这个对象中提取出来，等同于
    //const deploy = deployments.deploy;
    //const log = deployments.log
    const { deploy, log } = deployments;
    //getNameAccounts() 返回一个包含 deployer 属性的对象，等同于
    //const getNameAccountsResult = await getNameAccounts();
    //const deployer = getNameAccountsResult.deployer;
    log("Deploy Minatosys...");
    const { deployer } = await getNamedAccounts();
    const chainId = network.config.chainId;
  
    //---------------确认预言机地址-----------------
    if (developmentChains.includes(network.name)) {
      const ethUsdAggregator = await deployments.get("MockV3Aggregator");
      ethUsdPriceFeedAddress = ethUsdAggregator.address;
    } else {
      // ethUsdPriceFeedAddress的格式：0x694AA1769357215DE4FAC081bf1f309aDC325306
      ethUsdPriceFeedAddress = networkConfig[chainId]["ethUsdPriceFeed"];
    }
  
    // log(ethUsdPriceFeedAddress);
    //-----------------deploy-----------------------
    const args = [ethUsdPriceFeedAddress];
    const Minatosys= await deploy("Minatosys", {
      from: deployer,
      args: args, //喂价地址
      log: true,
      waitConfirmation: network.config.blockConfirmations || 1,
    });
    //-------------------verify-----------------------
    //当合约部署网络与指定的不符时会进行检查
    if (
      !developmentChains.includes(network.name) &&
      process.env.ETHERSCAN_API_KEY
    ) {
      //verify
      await verify(Minatosys.address, args);
    }
    log("--------------------------------------------------------");
  };
  
  module.exports.tags = ["all", "minatosys"];