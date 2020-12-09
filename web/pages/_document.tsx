import Document, { DocumentContext, Html, Head, Main, NextScript } from "next/document"
import { ServerStyleSheet } from "styled-components"
import { ServerStyleSheets } from "@material-ui/styles"
import { createMuiTheme, responsiveFontSizes } from "@material-ui/core/styles"
import React from "react"
import { CssBaseline } from "@material-ui/core"


const theme = responsiveFontSizes(createMuiTheme())

export default class MyDocument extends Document {
    render() {
        return (
            <Html>
                <Head>
                    <meta charSet="utf-8" />
                    <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no" />
                    <meta name="theme-color" content={theme.palette.primary.main} />
                    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Roboto+Slab:400,700|Material+Icons" />
                </Head>
                <body>
                    <Main />
                    <NextScript />
                </body>
            </Html>
        )
    }

    // static async getInitialProps(ctx: DocumentContext) {
    //     const sheets = new ServerStyleSheets()

    //     const originalRenderPage = ctx.renderPage

    //     ctx.renderPage = () =>
    //         originalRenderPage({
    //             enhanceApp: App => props => sheets.collect(<App {...props} />),
    //         })

        // const initialProps = await Document.getInitialProps(ctx)
    //     return {
            // ...initialProps,
    //         styles: (
    //             <React.Fragment key="styles">
    //                 {/* {initialProps.styles} */}
    //                 {sheets.getStyleElement()}
    //             </React.Fragment>
    //         ),
    //     }
    // }
}
