import { Button, Table, Jumbotron, Collapse } from "reactstrap";

const PurchaseDetail = (props) => {
	const onToggleItem = () => {
		props.toggleItem && props.toggleItem();
	};

	return (
		<>
			<Collapse isOpen={props.isOpen}>
				<Jumbotron className="purchase-detail-jumbotron">
					<Table className="purchase-table">
						<tbody>
							<tr>
								<th className="col-6">Name</th>
								<th className="col-2">Price</th>
								<th className="col-2">Quantity</th>
								<th className="col-2">Amount</th>
							</tr>
							{props.purchase.ProductPurchaseDetail.map((purchaseDetail) => {
								return (
									<tr>
										<td>{purchaseDetail.ProductName}</td>
										<td>{purchaseDetail.ProductPrice} RSD</td>
										<td>{purchaseDetail.ProductQuantity}</td>
										<td>
											{Number(purchaseDetail.ProductPrice) *
												Number(purchaseDetail.ProductQuantity)}{" "}
											RSD
										</td>
									</tr>
								);
							})}
						</tbody>
						<tfooter>
							<Button
								style={{ marginTop: "15px" }}
								className="my-button"
								type="button"
								onClick={onToggleItem}
							>
								Hide items
							</Button>
						</tfooter>
					</Table>
				</Jumbotron>
			</Collapse>
		</>
	);
};

export default PurchaseDetail;
