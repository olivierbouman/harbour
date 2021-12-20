import { useParams } from "react-router-dom"
import { useEffect, useState } from "react/cjs/react.development"
import styled from "styled-components"

const Boat = () => {
    const { id } = useParams()

    const Yellow = styled.div`
        padding: 25px;
        margin-left: 25px;
        width: 500px;
    `
    const [boat, setBoat] = useState({})

    useEffect(() => {
        setBoat({name: "Juliette", length: 880})
    }, [])    

    return (
        <Yellow>
            <table className="table table-compact table-striped">
                <thead></thead>
                <tbody>
                    <tr>
                        <td><strong>Title:</strong></td>
                        <td>{boat.name} - {id}</td>
                    </tr>
                    <tr>
                        <td><strong>Length:</strong></td>
                        <td>{boat.length}cm</td>
                    </tr>
                </tbody>
            </table>
        </Yellow>
    )
}

export default Boat