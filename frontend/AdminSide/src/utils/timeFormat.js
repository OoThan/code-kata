import moment from "moment";

export function dateFormat(dtime) {
  //   return moment(dtime).format("YYYY-MM-DD HH:mm:ss");
  return moment.parseZone(dtime).format("YYYY-MM-DD HH:mm:ss");
}

export function originalDate(dtime) {
  return moment.parseZone(dtime).format("YYYY-MM-DD ");
}

export function fromNowFormat(dtime) {
  return moment.parseZone(dtime).from();
}
