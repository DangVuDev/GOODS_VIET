import { Layout } from 'antd';
import ProductForm from './ProductForm';
import ProductList from './ProductList';

const { Header, Content } = Layout;

const Dashboard: React.FC = () => {

  
  return (
    <Layout className="min-h-screen">
      <Header className="bg-white text-white flex flex-col sm:flex-row items-center justify-between px-4 py-3 sm:py-0"
      style={{backgroundColor: "white"}}
      >
        <h1 className="text-lg sm:text-xl md:text-2xl font-bold mb-2 sm:mb-0 text-center sm:text-left">
          Quản lý sản phẩm
        </h1>
      </Header>

      <Content className="p-4 sm:p-6 md:p-8 bg-gray-50 min-h-[calc(100vh-64px)]">
        <div className="w-full max-w-5xl mx-auto space-y-6">
          <ProductForm />
          <ProductList />
        </div>
      </Content>
    </Layout>
  );
};

export default Dashboard;
