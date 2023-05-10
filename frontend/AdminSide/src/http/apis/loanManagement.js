import http from "../axios";

/**
 * @Loan_Packages
 */

export const getLoanPackageList = (data) =>
  http({
    url: "/loan-pkg/list",
    method: "POST",
    data,
  });

export const addLoanPackage = (data) =>
  http({
    url: "/loan-pkg/add",
    method: "POST",
    data,
  });

export const editLoanPackage = (data) =>
  http({
    url: "/loan-pkg/edit",
    method: "POST",
    data,
  });

export const deleteLoanPackage = (data) =>
  http({
    url: "/loan-pkg/delete",
    method: "POST",
    data,
  });
