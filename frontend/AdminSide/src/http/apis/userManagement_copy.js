import http from "../axios";

/**
 * @Teachers
 */

export const getTeachers = (data) =>
  http({ url: "/teacher/get-teachers", method: "GET", params: data });

export const addTeacher = (data) =>
  http({
    url: "/teacher/create",
    method: "POST",
    data,
  });

export const editTeacher = (data) =>
  http({
    url: "/teacher/update",
    method: "POST",
    data,
  });

export const deleteTeacher = (data) =>
  http({
    url: "/teacher/delete",
    method: "POST",
    data,
  });

/**
 * @Students
 */
export const getStudents = (data) =>
  http({ url: "/student/get-students", method: "GET", params: data });

export const addStudent = (data) =>
  http({
    url: "/student/create",
    method: "POST",
    data,
  });

export const editStudent = (data) =>
  http({
    url: "/student/update",
    method: "POST",
    data,
  });

export const deleteStudent = (data) =>
  http({
    url: "/student/delete",
    method: "POST",
    data,
  });

/**
 * @Student
 */

export const getStaffs = (data) =>
  http({ url: "/staff/get-staffs", method: "GET", params: data });

export const addStaff = (data) =>
  http({
    url: "/staff/create",
    method: "POST",
    data,
  });

export const editStaff = (data) =>
  http({
    url: "/staff/update",
    method: "POST",
    data,
  });

export const deleteStaff = (data) =>
  http({
    url: "/staff/delete",
    method: "POST",
    data,
  });
