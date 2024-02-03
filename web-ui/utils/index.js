import { useState } from "react"

export function getUserInfo() {
  const [uerInfo, _setUerInfo] = useState({
    userId: 5
  });

  return uerInfo
}
