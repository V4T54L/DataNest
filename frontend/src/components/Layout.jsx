import { Outlet } from 'react-router-dom';

const Layout = ({ children }) => {
    const handleLogout = () => {
        // Add your logout logic here (e.g., clearing tokens, redirecting to login page)
        console.log("Logged out");
    };

    return (
        <div className="flex flex-col h-screen bg-gray-50">
            {/* Header */}
            <header className="flex items-center justify-between p-4 bg-white shadow">
                <h1 className="text-xl font-bold">Dashboard Home</h1>
                <div className="flex items-center">
                    <div className="bg-gray-300 rounded-full w-8 h-8 flex items-center justify-center mr-2">
                        <span className="text-lg">👤</span>
                    </div>
                    <button
                        onClick={handleLogout}
                        className="bg-red-500 text-white py-1 px-3 rounded hover:bg-red-600 transition"
                    >
                        Logout
                    </button>
                </div>
            </header>

            {/* Main Section */}
            <main className="flex-1 p-4">
                <Outlet />
            </main>

            {/* Footer */}
            <footer className="bg-gray-200 text-center py-4">
                <p>&copy; {new Date().getFullYear()} Your Company Name. All rights reserved.</p>
            </footer>
        </div>
    );
};

export default Layout