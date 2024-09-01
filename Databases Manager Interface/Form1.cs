using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Net.Http;
using System.Reflection;
using System.Security.Policy;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using static System.Windows.Forms.VisualStyles.VisualStyleElement;

namespace Databases_Manager_Interface
{
     
    public partial class Form1 : Form
    {

        public Form1()
        {
            Program.f1 = this;
            InitializeComponent();
            costyl("localhost");

        }

        public async void costyl(string hostname)
        {
            string url = "http://" + hostname + ":8080/initDBlist";
            string responseContent = await SendGetRequestAsync(url);
            //dataGridView1.CellClick += new DataGridViewCellEventHandler(dataGridView1_CellContentClick);
            
            string[] statuses = responseContent.Split(new char[] { '|' }, StringSplitOptions.RemoveEmptyEntries);
            for(int i= 0; i<statuses.Length; i++)
            {
                Program.f1.dataGridView1.Rows.Add(statuses[i]);
            }
        }

        private void Form1_Load(object sender, EventArgs e)
        {
        }

        private void textBox1_TextChanged(object sender, EventArgs e)
        {
        }

        private void button1_Click(object sender, EventArgs e)
        {
            AddForm addfrm = new AddForm();
            addfrm.Show();
        }

        private async void dataGridView1_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {
            
               string dbname = dataGridView1[e.ColumnIndex, e.RowIndex].Value.ToString();
               //if (dbname == "") { MessageBox.Show("Welcome to Postgres Database Manager Interface"); }
            
               string data = "{\"database_name\":\"" + dbname + "\"}";

            try
            {
                ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
                string url = "http://localhost:8080/workDB";
                ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
                int k = 0;
                int index = 0;
                int j = 0;
                string[] ss=new string[1000];
                List<string> s = new List<string>();

                
                CustomResponse response = await GetApiResponseAsync(url, data);
                int tablesCounter = response.Tables.Count;


                foreach (var table in response.Tables)
                {
                    index++;
                    
                    Form4 frm4 = new Form4();
                    frm4.Text = $"DB: {response.DatabaseName} / Table: {table.TableName} / Index: {index}";
                    frm4.Show();
                    
                    foreach (var row in table.Rows)
                    {
                        
                        foreach (var kvp in row)
                        {
                            if (j < 1)
                            {
                                frm4.dataGridView1.Columns.Add($"{kvp.Key}", $"{kvp.Key}");
                            }
                                ss[k] = $"{kvp.Value}";
                                k++;
                            
                        }
                        
                         
                        frm4.dataGridView1.Rows.Add(ss);
                        

                        //}
                        for (int i = 0; i <= k; i++)
                        {
                           ss[i] = "";
                        }
                        k = 0;
                        j++;
                    }
                    
                    j = 0;
                    int lastRowIndex = frm4.dataGridView1.Rows.Count - 1;
                    // Извлекаем последнюю строку
                    DataGridViewRow lastRow = frm4.dataGridView1.Rows[lastRowIndex];
                    lastRow.DefaultCellStyle.BackColor = Color.Red;
                }       
            }
            catch (Exception ex)
            {
                Form4 frm4 = new Form4();
                frm4.Text = $"DB: {dbname} / Table:  / Index:  1";
                frm4.Show();
                MessageBox.Show($"{dbname} database is empty");
            }
        }

        static async Task<string> SendGetRequestAsync(string url)
        {
            using (HttpClient client = new HttpClient())
            {
                try
                {
                       
                    HttpResponseMessage response = await client.GetAsync(url);

                    // Проверка успешности запроса
                    response.EnsureSuccessStatusCode();

                    // Чтение и возврат содержимого ответа
                    string responseContent = await response.Content.ReadAsStringAsync();
                    return responseContent;
                }
                catch (HttpRequestException e)
                {
                    Console.WriteLine($"Ошибка при отправке запроса: {e.Message}");
                    return null;
                }
            }
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

        public static async Task<CustomResponse> GetApiResponseAsync(string url, string data)
        {
            string jsonString = await FetchDataAsync(url, data);
            return JsonConvert.DeserializeObject<CustomResponse>(jsonString);
        }

        private void button1_Click_1(object sender, EventArgs e)
        {
            MessageBox.Show("You also can create a new database by sql-request in your connected root database");
            Form7 form7 = new Form7();
            form7.Show();
        }

        private void button2_Click(object sender, EventArgs e)
        {
            Program.f1.dataGridView1.Rows.Clear();
            costyl("localhost");
        }
    }
}
