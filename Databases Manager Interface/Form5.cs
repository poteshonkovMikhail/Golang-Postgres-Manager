using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Net.Http;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Xml.Linq;
using static System.Net.Mime.MediaTypeNames;
using static System.Windows.Forms.VisualStyles.VisualStyleElement;

namespace Databases_Manager_Interface
{
    public partial class Form5 : Form
    {
        public Form5()
        {
            InitializeComponent();
        }

        public async void button1_Click(object sender, EventArgs e)
        {
            string url1 = "http://localhost:8080/db-info";
            string url2 = "http://localhost:8080/sql-request";
            string dbname = ExtractTextAfterColons(this.Text);
            //MessageBox.Show(dbname);
            string data = "{\"database_name\":\"" + dbname + "\"}";
           
            DbInfo response = await GetApiResponseAsync(url1, data);
            string richtext = richTextBox1.Text.Replace("\n", string.Empty);
            string dataTwo = "{\"username_parameter\":\"" + response.DatabaseUsername + "\"," +
            " \"password_parameter\":\"" + response.DatabasePassword + "\"," +
            " \"hostname_parameter\":\"" + response.DatabaseHostname + "\"," +
            " \"database_name_parameter\":\"" + response.DatabaseName + "\"," +
            " \"port_parameter\":\"" + "5432" + "\"," +
                " \"sqlCommand_parameter\":\"" + richtext + "\"}";

            SqlResponse sqlResponse = await SendAsyncSqlRequest(url2, dataTwo);
            int k = 0;
            string[] ss = new string[1000];
            Form6 frm6 = new Form6();
            frm6.Text = richTextBox1.Text;
            frm6.Show();
            foreach (var row in sqlResponse.Results)
            {
                foreach (var kvp in row)
                {
                    frm6.dataGridView1.Columns.Add($"{kvp.Key}", $"{kvp.Key}");
                    ss[k] = $"{kvp.Value}";
                    k++;
                }
                frm6.dataGridView1.Rows.Add(ss);

                for (int i = 0; i <= k; i++)
                {
                    ss[i] = "";
                }
                k = 0;
            }

            int lastRowIndex = frm6.dataGridView1.Rows.Count - 1;
            // Извлекаем последнюю строку
            DataGridViewRow lastRow = frm6.dataGridView1.Rows[lastRowIndex];
            lastRow.DefaultCellStyle.BackColor = Color.HotPink;

            //MessageBox.Show(sqlResponse.Results);

        }

        static string ExtractTextAfterColons(string input)
        {
            int startIndex = input.IndexOf(':') + 1; // Позиция символа : + 1, чтобы начать после :
            //int endIndex = input.IndexOf('/', startIndex); // Позиция символа / после :

            // Если символы : или / не найдены, вернуть пустую строку
            

            // Извлечение и обрезка пробелов по краям
            string result = input.Substring(startIndex).Trim();
            return result;
        }

        private static async Task<string> FetchDataAsync(string url, string data)
        {
            using (HttpClient client = new HttpClient())
            {
                var content = new StringContent(data, Encoding.UTF8, "application/json");
                HttpResponseMessage response = await client.PostAsync(url, content);
                response.EnsureSuccessStatusCode();
                return await response.Content.ReadAsStringAsync();
            }
        }

        public static async Task<DbInfo> GetApiResponseAsync(string url, string data)
        {
            string jsonString = await FetchDataAsync(url, data);
            return JsonConvert.DeserializeObject<DbInfo>(jsonString);
        }

        public static async Task<SqlResponse> SendAsyncSqlRequest(string url, string data)
        {
            string jsonString = await FetchDataAsync(url, data);
            return JsonConvert.DeserializeObject<SqlResponse>(jsonString);
        }
    }
}
