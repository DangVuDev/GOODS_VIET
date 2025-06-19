
import { Card, QRCode } from 'antd';

interface QRCodeDisplayProps {
  qrCodeUrl: string;
}

const QRCodeDisplay: React.FC<QRCodeDisplayProps> = ({ qrCodeUrl }) => {
  return (
    <Card className="mt-4 text-center card-hover">
      <h3 className="text-lg font-semibold text-blue-800">Mã QR sản phẩm</h3>
      <div className="flex justify-center my-4">
        <QRCode value={qrCodeUrl} size={200} className="rounded-md" />
      </div>
      <p className="text-gray-600">Quét mã để xem chi tiết sản phẩm</p>
    </Card>
  );
};

export default QRCodeDisplay;