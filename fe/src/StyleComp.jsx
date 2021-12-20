import styled from "styled-components"

const Div = styled.div`
    text-align: center;
    margin: 40px;
    border: 5px outset blue;
    &:hover {
        background-color: orange;
    };
`


const Stylerr = () => {
    return (
        <Div>
            Lets do some stylingh
        </Div>
    )
}



export default Stylerr
