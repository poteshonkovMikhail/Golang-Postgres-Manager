using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Databases_Manager_Interface
{
    public partial class Form4 : Form
    {
        
        public Form4()
        {
            InitializeComponent();
            
        }

        private void dataGridView1_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            string st = ExtractTextBetweenColonsAndSlashes(this.Text);
            Form5 form5 = new Form5();
            form5.Text="SQL Request: " + st;
            form5.Show(); 
        }

        static string ExtractTextBetweenColonsAndSlashes(string input)
        {
            int startIndex = input.IndexOf(':') + 1; // Позиция символа : + 1, чтобы начать после :
            int endIndex = input.IndexOf('/', startIndex); // Позиция символа / после :

            // Если символы : или / не найдены, вернуть пустую строку
            if (startIndex < 1 || endIndex == -1)
                return string.Empty;

            // Извлечение и обрезка пробелов по краям
            string result = input.Substring(startIndex, endIndex - startIndex).Trim();
            return result;
        }
    }
}
