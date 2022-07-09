import { useState, useEffect } from "react";
import { Card, Col, Table } from "reactstrap";
import axios from "axios";
import "../../Assets/css/product-detail.css";
import * as videoGameAPI from "../../APIs/ProductMicroservice/video_game_api";

const VideoGameDetail = (props) => {
	const [product, setProduct] = useState(null);

	useEffect(() => {
		getById();
	}, []);

	const getById = () => {
		axios
			.get(`${videoGameAPI.GET_BY_ID}/${props.product.Id}`)
			.then((res) => {
				setProduct(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	return (
		<Col style={{ paddingTop: "5px", paddingBottom: "10px" }} md={7}>
			<Card className="product-detail-table-card">
				<Table className="product-detail-table">
					{product !== null &&
						Object.keys(product).map(function (value, idx) {
							if (value !== "Product") {
								if (typeof product[value] == "boolean") {
									return (
										<tr key={idx}>
											<th>{value}</th>
											<td>{product[value] ? "Yes" : "No"}</td>
										</tr>
									);
								} else {
									return (
										<tr key={idx}>
											<th>{value}</th>
											<td>{product[value].toString()}</td>
										</tr>
									);
								}
							}
						})}
				</Table>
			</Card>
		</Col>
	);
};

export default VideoGameDetail;
