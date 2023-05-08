export function checkPermissions(permission) {
  // console.log("Here checnk permission");
  let permissionsList = JSON.parse(localStorage.getItem("permissionList"));
  let permissionIds;

  if (permissionsList) {
    permissionIds = permissionsList.map((item) => {
      return item.id;
    });
  }

  // console.log(permissionIds);
  if (permissionsList) {
    // console.log(permissionIds.includes(permission));
    return permissionIds.includes(permission);
  }
}
