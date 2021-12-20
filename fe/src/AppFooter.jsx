import { Fragment } from "react/cjs/react.production.min"

const Footer = ({name, counter}) => {
    const currentYear = new Date().getFullYear()
    return (
        <Fragment>
            <hr />
            <p>Copyright &copy; {currentYear} Bouman & {name} Ltd. - {counter}</p>
        </Fragment>
    )
}

export default Footer
