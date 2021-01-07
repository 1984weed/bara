import { useSession } from 'next-auth/client';
import { useRouter } from "next/router";
import { useEffect } from "react";


// It redirects to top page for not admin users. Admin page should be allowed only admin user.
export function useProtectAdminPage(
): void {
    const router = useRouter()
    const [session] = useSession()

    useEffect(() => {
        if (session == null || session.user.role !== "admin") {
            router.push("/")
        }
    }, [session])
}
