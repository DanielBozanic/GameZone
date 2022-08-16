import { useState, useEffect } from "react";
import {
	CardBody,
	Card,
	Row,
	Col,
	Input,
	Spinner,
	CardTitle,
} from "reactstrap";
import { Helmet } from "react-helmet";
import { Chart, registerables } from "chart.js";
import { Bar } from "react-chartjs-2";
import axios from "axios";
import "../Assets/css/business-reports.css";
import * as businessReportAPI from "../APIs/BusinessReportMicroservice/business_report_api";
Chart.register(...registerables);

const BusinessReports = () => {
	const [reportUrl, setReportUrl] = useState(
		businessReportAPI.GET_PRODUCTS_WITH_BIGGEST_PROFIT_LAST_THIRTY_DAYS
	);
	const [reportData, setReportData] = useState(null);
	const [title, setTitle] = useState(
		"Products with the biggest profit in last 30 days"
	);
	const [loading, setLoading] = useState(true);

	const options = {
		responsive: true,
		plugins: {
			legend: {
				display: false,
			},
			title: {
				display: true,
				text: title,
				color: "white",
			},
			tooltip: {
				callbacks: {
					title: (context) => {
						return context[0].label.replaceAll(",", " ");
					},
				},
			},
		},
		scales: {
			y: {
				ticks: { color: "white" },
			},
			x: {
				ticks: {
					color: "white",
					maxRotation: 90,
					minRotation: 0,
					callback: function (value) {
						return this.getLabelForValue(value).length > 10
							? this.getLabelForValue(value).substr(0, 10) + "..."
							: this.getLabelForValue(value);
					},
				},
			},
		},
		maintainAspectRatio: false,
	};

	useEffect(() => {
		getReport();
	}, [reportUrl]);

	const getReport = () => {
		axios
			.get(`${reportUrl}`)
			.then((res) => {
				setReportData({
					labels: Object.keys(res.data),
					datasets: [
						{
							data: Object.values(res.data),
							backgroundColor: "orange",
							borderColor: "white",
						},
					],
				});
				setLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const setupReport = (e) => {
		setLoading(true);
		const index = e.nativeEvent.target.selectedIndex;
		setTitle(e.nativeEvent.target[index].text);
		setReportUrl(e.target.value);
	};

	return (
		<>
			<Helmet>
				<title>Business Reports | GameZone</title>
			</Helmet>
			<Card className="card business-report-card">
				<CardBody>
					<Row>
						<Col>
							<Input
								className="input-field report-select"
								type="select"
								onChange={(e) => setupReport(e)}
							>
								<option hidden>Select business report</option>
								<option
									value={
										businessReportAPI.GET_PRODUCTS_WITH_BIGGEST_PROFIT_LAST_THIRTY_DAYS
									}
								>
									Products with the biggest profit in last 30 days
								</option>
								<option
									value={businessReportAPI.GET_SOLD_VIDEO_GAMES_BY_PLATFORM}
								>
									Sold video games by platform
								</option>
								<option value={businessReportAPI.GET_SOLD_VIDEO_GAMES_BY_FORM}>
									Sold video games by form
								</option>
							</Input>
						</Col>
					</Row>
					{loading && (
						<div className="div-spinner">
							<Spinner className="spinner" />
						</div>
					)}
					{!loading && reportData !== null && (
						<>
							{reportData.datasets[0].data.length > 0 && (
								<Row>
									<Col>
										<Bar
											className="report-chart"
											options={options}
											data={reportData}
										/>
									</Col>
								</Row>
							)}
							{reportData.datasets[0].data.length <= 0 && (
								<Row>
									<Col>
										<CardTitle className="title" tag="h5">
											No data available
										</CardTitle>
									</Col>
								</Row>
							)}
						</>
					)}
				</CardBody>
			</Card>
		</>
	);
};

export default BusinessReports;
