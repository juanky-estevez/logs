export let utcOffset = 0;

export const defineUTC = (): void => {
  const timezone = process.env.TIMEZONE;
  if (!timezone || timezone === "" || !timezone.includes("UTC")) return;

  const offsetString = timezone.replace("UTC", "");
  if (offsetString === "") return;

  const offsetNumber = Number(offsetString);
  if (isNaN(offsetNumber)) return;

  utcOffset = offsetNumber;
};