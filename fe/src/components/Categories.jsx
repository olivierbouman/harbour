
import { Link, useMatch } from 'react-router-dom'

const Categories = () => {
    const { pathnameBase } = useMatch('/by-category')

    return (
        <div>
            <h2>CATTTTTTegories</h2>
            <ul>
                <li><Link to={`${pathnameBase}/yachts`}>Sailing Yachts</Link></li>
                <li><Link to={`${pathnameBase}/motorboats`}>Motor Boats</Link></li>
                <li><Link to={`${pathnameBase}/other-categories`}>Other categories</Link></li>

            </ul>
        </div>
        
    )
}

export default Categories