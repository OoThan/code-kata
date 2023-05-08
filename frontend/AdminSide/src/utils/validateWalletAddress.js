import Web3 from "web3";
var web3 = new Web3(Web3.givenProvider);

export const validateWalletAddress = () => {
  return (rule, value, callback) => {
    if (state.form.address_type == "ETH") {
      let isAddETH = web3.utils.isAddress(value);
      if (!isAddETH) {
        callback(new Error(t("common.notValidateETHAdd")));
      } else {
        callback();
      }
    } else {
      let isAddTRX = tronWeb.isAddress(value);
      if (!isAddTRX) {
        callback(new Error(t("common.notValidateTRXAdd")));
      } else {
        callback();
      }
    }
  };
};
