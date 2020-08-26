import LinkNext from "next/link"
import { Link } from "@material-ui/core"

type Props = {
    href: string
    label: string
}

export const MaterialLink = ({ href, label }: Props) => {
    return (
        <LinkNext href={href}>
            <Link href="#" variant="subtitle1" color="textSecondary">
                {label}
            </Link>
        </LinkNext>
    )
}
