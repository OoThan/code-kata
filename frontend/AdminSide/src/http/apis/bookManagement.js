import http from "../axios";

/**
 * @Category
 */

export const getCategoryList = (data) =>
  http({ url: "/book-category/get-categories", method: "GET", params: data });

export const addCategoryList = (data) =>
  http({
    url: "/book-category/create",
    method: "POST",
    data,
  });

export const editCategoryList = (data) =>
  http({
    url: "/book-category/update",
    method: "POST",
    data,
  });

export const deleteCategoryList = (data) =>
  http({
    url: "/book-category/delete",
    method: "POST",
    data,
  });

/**
 * @SubCategory
 */

export const getSubCategoryList = (data) =>
  http({
    url: "/book-sub-category/get-categories",
    method: "GET",
    params: data,
  });

export const addSubCategoryList = (data) =>
  http({
    url: "/book-sub-category/create",
    method: "POST",
    data,
  });

export const editSubCategoryList = (data) =>
  http({
    url: "/book-sub-category/update",
    method: "POST",
    data,
  });

export const deleteSubCategoryList = (data) =>
  http({
    url: "/book-sub-category/delete",
    method: "POST",
    data,
  });

/**
 * @Books
 */
export const getBookList = (data) =>
  http({ url: "/book/get-books", method: "GET", params: data });

export const addBookList = (data) =>
  http({
    url: "/book/create",
    method: "POST",
    data,
  });

export const editBookList = (data) =>
  http({
    url: "/book/update",
    method: "POST",
    data,
  });

export const deleteBookList = (data) =>
  http({
    url: "/book/delete",
    method: "POST",
    data,
  });
