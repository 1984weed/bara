// import { useState, useEffect, useContext, createContext, createElement, ReactNode } from "react"

// const SessionContext = createContext(undefined)

// // Client side method
// export const useSession = (session?: any): any => {
//     // Try to use context if we can
//     const value = useContext(SessionContext)

//     // If we have no Provider in the tree, call the actual hook
//     if (value == null) {
//         return _useSessionHook(session)
//     }

//     return value
// }

// const _useSessionHook = (session): any[] => {
//     const [data, setData] = useState(session)
//     const [loading, setLoading] = useState(true)
//     const _getSession = async () => {
//         try {
//             const newClientSessionData = await getSession()
//             if(newClientSessionData) {
//                 setData(newClientSessionData)
//                 localStorage.setItem("api-token", newClientSessionData.token)
//                 setLoading(false)
//             }
//         } catch (error) {
//             console.log("client session error", error)
//         }
//     }
//     useEffect(() => {
//         _getSession()
//     }, [])

//     return [data, loading]
// }

// function getSession(): Promise<any> {
//     return new Promise((resolve, reject) => {
//         const baseUrl = "http://localhost:3000" 

//         fetch(`${baseUrl}/auth/session`, {
//             credentials: "same-origin",
//         })
//             .then(res => {
//                 if (res.ok) {
//                     return res.json()
//                 } else {
//                     reject()
//                 }
//             })
//             .then(res => resolve(res))
//     })
// }

// export const Provider = ({ children, session }: { children: ReactNode; session: any }) => {
//     return createElement(SessionContext.Provider, { value: useSession(session) }, children)
// }
