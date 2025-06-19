import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Table, Card, type Breakpoint } from 'antd';
import { getAllProducts } from '../services/api';
import type { Product } from '../types';

const ProductList: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const data = await getAllProducts();
        setProducts(data);
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    };
    fetchProducts();
  }, []);

  const columns = [
    { title: 'ID', dataIndex: 'productId', key: 'productId', responsive: ['md' as Breakpoint] },
    { title: 'Tên sản phẩm', dataIndex: 'name', key: 'name' },
    { title: 'Nhà sản xuất', dataIndex: 'manufacturer', key: 'manufacturer', responsive: ['md' as Breakpoint] },
    {
      title: 'Trạng thái mới nhất',
      key: 'latestStatus',
      render: (record: Product) => record.statuses[record.statuses.length - 1]?.status,
    },
  ];

  return (
    <Card className="card-hover" title={<h2 className="text-xl font-bold">Danh sách sản phẩm</h2>}>
      <Table
        dataSource={products}
        columns={columns}
        rowKey="productId"
        pagination={{ pageSize: 5 }}
        scroll={{ x: true }}
        onRow={(record: Product) => ({
          onClick: () => navigate(`/product/${record.productId}`),
          className: 'cursor-pointer hover:bg-gray-50',
        })}
      />
    </Card>
  );
};

export default ProductList;