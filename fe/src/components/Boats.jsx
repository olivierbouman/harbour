import { Fragment, useState, useEffect } from "react"
import { Link, Outlet } from "react-router-dom"
import styled from "styled-components"

const Boats = () => {

    const [boats, setBoats] = useState([])

    useEffect(() => {
        setBoats(
            [
                {id: 1, name: "Juliette", length: 880},
                {id: 2, name: "King Fisher", length: 760},
                {id: 3, name: "De Bak", length: 2100},
                {id: 4, name: "Tit sal et leren", length: 2100},
            ]
        )
    }, [])


    const FlexBox = styled.div`
        display: flex;
        flex-direction: row;
    `

    return (
        <Fragment>
            <h2>Look at them boats</h2>
            <FlexBox>
                <ul>
                    {boats.map((boat) => (
                        <li key={boat.id}>
                            <Link to={`${boat.id}`}>
                                {boat.name} - {boat.length}cm
                            </Link>
                        </li>
                    ))}
                </ul>
                <Outlet />
            </FlexBox>
        </Fragment>
    )
}

export default Boats