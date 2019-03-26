module.exports = {
    up: (queryInterface, Sequelize) => {
        return queryInterface.addColumn('Comment', 'path', {
                type: Sequelize.STRING
            });
    },
    down: (queryInterface, Sequelize) => {
        return
    }
};
