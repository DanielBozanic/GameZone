import AppNavbar from "../Layout/AppNavbar";
import ProductsView from "../Components/ProductsView";
import Search from "../Components/Search";
import { Row, Container } from "reactstrap";
import axios from "axios";
import { useState, useEffect } from "react";
import { Helmet } from "react-helmet";

const ProductList = (props) => {
	const [products, setProducts] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [numberOfRecords, setNumberOfRecords] = useState(0);
	const [filter, setFilter] = useState(null);
	const [showFilter, setShowFilter] = useState(true);
	const [searchTerm, setSearchTerm] = useState("");
	const [clearSearchTerm, setClearSearchTerm] = useState(false);
	const pageSize = 12;

	const handleClick = (index) => {
		setCurrentPage(index);
	};

	useEffect(() => {
		const pageSetup = () => {
			if (filter !== null) {
				getProductsFilter();
				getNumberOfRecordsFilter();
			} else if (searchTerm !== "") {
				getProductsSearch();
				getNumberOfRecordsSearch();
			} else {
				getProducts();
				getNumberOfRecords();
				setShowFilter(true);
			}
		};
		pageSetup();
	}, [currentPage, filter, searchTerm]);

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

	const getNumberOfRecords = () => {
		axios
			.get(`${props.GET_NUMBER_OF_RECORDS}`)
			.then((res) => {
				setNumberOfRecords(res.data);
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

	const getNumberOfRecordsFilter = () => {
		axios
			.post(`${props.GET_NUMBER_OF_RECORDS_FILTER}`, filter)
			.then((res) => {
				setNumberOfRecords(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getProductsSearch = () => {
		axios
			.get(
				`${props.SEARCH_BY_NAME}?page=${currentPage}&pageSize=${pageSize}&name=${searchTerm}`
			)
			.then((res) => {
				setProducts(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getNumberOfRecordsSearch = () => {
		axios
			.get(`${props.GET_NUMBER_OF_RECORDS_SEARCH}?name=${searchTerm}`)
			.then((res) => {
				setNumberOfRecords(res.data);
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
		if (st !== "") {
			setFilter(null);
			setShowFilter(false);
		}
		setCurrentPage(1);
	};

	return (
		<>
			<Helmet>
				<title>{props.title} | GameZone</title>
			</Helmet>
			<Container>
				<Row style={{ display: "flex" }}>
					{showFilter && props.filter && (
						<props.filter onFilterClick={onFilterClick} />
					)}
					<Search
						onSearchClick={onSearchClick}
						clearSearchTerm={clearSearchTerm}
						searchPlaceholder={props.searchPlaceholder}
					/>
				</Row>
			</Container>
			<AppNavbar />
			<ProductsView
				products={products}
				currentPage={currentPage}
				pageSize={pageSize}
				numberOfRecords={numberOfRecords}
				handleClick={handleClick}
			/>
		</>
	);
};

export default ProductList;
