import { useState, useEffect } from "react";
import { CardBody, Card, Row, Col, Input } from "reactstrap";
import { Chart, registerables } from "chart.js";
import { Bar } from "react-chartjs-2";
import axios from "axios";
import "../Assets/css/business-reports.css";
import * as businessReportAPI from "../APIs/BusinessReportMicroservice/business_report_api";
Chart.register(...registerables);

const BusinessReports = () => {
	const [reportUrl, setReportUrl] = useState(
		businessReportAPI.GET_PRODCUTS_WITH_BIGGEST_PROFIT_LAST_THIRTY_DAYS
	);
	const [reportData, setReportData] = useState(null);
	const [title, setTitle] = useState(
		"Products with the biggest profit in last 30 days"
	);

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
		},
		scales: {
			y: {
				ticks: { color: "white" },
			},
			x: {
				ticks: { color: "white" },
			},
		},
		maintainAspectRatio: false,
	};

	useEffect(() => {
		getReport();
	}, [reportUrl]);

	const getReport = () => {
		const instance = axios.create();
		delete instance.defaults.headers.common["Authorization"];
		instance
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
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const setupReport = (e) => {
		const index = e.nativeEvent.target.selectedIndex;
		setTitle(e.nativeEvent.target[index].text);
		setReportUrl(e.target.value);
	};

	return (
		<>
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
										businessReportAPI.GET_PRODCUTS_WITH_BIGGEST_PROFIT_LAST_THIRTY_DAYS
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
					{reportData !== null && (
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
				</CardBody>
			</Card>
		</>
	);
};

export default BusinessReports;
