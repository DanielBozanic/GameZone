import * as productAPI from "../APIs/ProductMicroservice/product_api";
import AppNavbar from "../Layout/AppNavbar";
import ProductView from "../Components/ProductView";
import Filter from "../Components/VideoGamesFilter";
import Search from "../Components/Search";
import { Row, Container } from "reactstrap";
import axios from "axios";
import { useState, useEffect } from "react";

const ProductList = (props) => {
	const [products, setProducts] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [pageCount, setPageCount] = useState([]);
	const [filter, setFilter] = useState(null);
	const [showFilter, setShowFilter] = useState(true);
	const [searchTerm, setSearchTerm] = useState("");
	const [clearSearchTerm, setClearSearchTerm] = useState(false);
	const pageSize = 8;

	const handleClick = (e, index) => {
		e.preventDefault();
		setCurrentPage(index);
	};

	useEffect(() => {
		pageSetup();
	}, [currentPage, filter, searchTerm]);

	const pageSetup = () => {
		if (filter !== null) {
			getProductsFilter();
			getPageCountFilter();
		} else if (searchTerm !== "") {
			getProductsSearch();
			getPageCountSearch();
		} else {
			getProducts();
			getPageCount();
			setShowFilter(true);
		}
	};

	const getProducts = () => {
		axios
			.get(`${props.GET_PRODUCTS}?page=${currentPage}&pageSize=${pageSize}`)
			.then((res) => {
				setProducts(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPageCount = () => {
		axios
			.get(`${props.GET_NUMBER_OF_RECORDS}`)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data) / pageSize));
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getProductsFilter = () => {
		axios
			.post(`${props.FILTER}?page=${currentPage}&pageSize=${pageSize}`, filter)
			.then((res) => {
				setProducts(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPageCountFilter = () => {
		axios
			.post(`${props.GET_NUMBER_OF_RECORDS_FILTER}`, filter)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data) / pageSize));
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getProductsSearch = () => {
		axios
			.get(
				`${productAPI.SEARCH_BY_NAME}?page=${currentPage}&pageSize=${pageSize}&name=${searchTerm}`
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

	const onFilterClick = (filterData) => {
		setFilter(filterData);
		setSearchTerm("");
		setClearSearchTerm(!clearSearchTerm);
		setCurrentPage(1);
	};

	const onSearchClick = (st) => {
		setSearchTerm(st);
		setFilter(null);
		setShowFilter(false);
		setCurrentPage(1);
	};

	return (
		<>
			<Container>
				<Row style={{ display: "flex" }}>
					{showFilter && <Filter onFilterClick={onFilterClick} />}
					<Search
						onSearchClick={onSearchClick}
						clearSearchTerm={clearSearchTerm}
					/>
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

export default ProductList;
