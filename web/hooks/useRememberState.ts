import React, { Dispatch } from "react";

export function useRememberState(
  key: string,
  initialValue: string = ""
): [string, Dispatch<string>] {
  const [value, setValue] = React.useState(() => {
    if (typeof window !== "undefined") {
      if(localStorage.getItem(key) == null || localStorage.getItem(key) === "") {
        return initialValue
      }
      return localStorage.getItem(key)
    }
    return initialValue;
  });

  React.useEffect(() => {
    localStorage.setItem(key, value);
  }, [value]);

  return [value, setValue];
}
