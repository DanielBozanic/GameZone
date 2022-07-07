import * as productAPI from "../APIs/ProductMicroservice/product_api";
import AppNavbar from "../Layout/AppNavbar";
import Search from "../Components/Search";
import ProductView from "../Components/ProductView";
import { Row, Container } from "reactstrap";
import axios from "axios";
import { useState, useEffect } from "react";

const Main = () => {
	const [products, setProducts] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [pageCount, setPageCount] = useState([]);
	const [searchTerm, setSearchTerm] = useState("");
	const pageSize = 8;

	const handleClick = (e, index) => {
		e.preventDefault();
		setCurrentPage(index);
	};

	useEffect(() => {
		if (searchTerm !== "") {
			getProductsSearch();
			getPageCountSearch();
		} else {
			setProducts([]);
		}
	}, [currentPage, searchTerm]);

	const getProductsSearch = () => {
		axios
			.get(
				`${productAPI.SEARCH_BY_NAME}
					?page=${currentPage}&pageSize=${pageSize}&name=${searchTerm}`
			)
			.then((res) => {
				setProducts(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPageCountSearch = () => {
		axios
			.get(`${productAPI.GET_NUMBER_OF_RECORDS_SEARCH}?name=${searchTerm}`)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data) / pageSize));
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const onSearchClick = (st) => {
		setSearchTerm(st);
		setCurrentPage(1);
	};

	return (
		<>
			<Container>
				<Row style={{ display: "flex" }}>
					<Search onSearchClick={onSearchClick} />
				</Row>
			</Container>
			<AppNavbar />
			<ProductView
				products={products}
				currentPage={currentPage}
				pageCount={pageCount}
				handleClick={handleClick}
			/>
		</>
	);
};

export default Main;
