import "../Assets/css/filter.css";
import { Button, Row, Col, Input } from "reactstrap";
import { useState, useEffect } from "react";

const Search = (props) => {
	const [searchTerm, setSearchTerm] = useState();

	useEffect(() => {
		setSearchTerm("");
	}, [props.clearSearchTerm]);

	const search = () => {
		props.onSearchClick && props.onSearchClick(searchTerm);
	};

	return (
		<>
			<Col md="10">
				<Row>
					<Col md="9" style={{ textAlign: "right" }}>
						<Input
							className="input-field search-bar"
							name="searchTerm"
							type="text"
							value={searchTerm}
							placeholder={props.searchPlaceholder}
							onChange={(e) => setSearchTerm(e.target.value)}
						/>
					</Col>
					<Col md="3">
						<Button className="my-button search-btn" onClick={search}>
							Search
						</Button>
					</Col>
				</Row>
			</Col>
		</>
	);
};

export default Search;
